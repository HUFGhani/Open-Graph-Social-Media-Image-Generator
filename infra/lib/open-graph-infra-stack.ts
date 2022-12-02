import * as cdk from 'aws-cdk-lib';
import {
    aws_apigateway as apigw,
    aws_certificatemanager as certificateManager,
    aws_cloudfront as cloudfront,
    aws_lambda as lambda,
    aws_route53 as route53,
    aws_route53_targets,
    aws_s3 as s3,
    aws_s3_assets as asserts,
    aws_ssm as ssm,
  RemovalPolicy
} from 'aws-cdk-lib';
import {Construct} from 'constructs';
import * as path from "path";

import {CloudFrontToApiGateway} from "@aws-solutions-constructs/aws-cloudfront-apigateway";



export class OpenGraphInfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

  const lambdaAsset = new asserts.Asset(this, "HelloGoServerLambdaFnZip",{
    path: path.join(__dirname,"../../functions")
  })

    const opengraphAsset = new s3.Bucket(this,'opengraphAsset',{
      removalPolicy: RemovalPolicy.DESTROY,
    }
  )

    const lambdaFN = new lambda.Function(this,"HelloGoServerLambdaFn",{
      code:lambda.Code.fromBucket(lambdaAsset.bucket,lambdaAsset.s3ObjectKey),
      // timeout: Duration.seconds(300),
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
      memorySize: 521,
      environment: {
        S3_BUCKET_NAME: opengraphAsset.bucketName,
      },

    })
    opengraphAsset.grantReadWrite(lambdaFN)

   const apiGateway = new apigw.LambdaRestApi(
        this, "LambdaFnEndpoint",{
          handler: lambdaFN,
          binaryMediaTypes:['*/*']
        }
    )

    const distribution = new CloudFrontToApiGateway(this, 'test-cloudfront-apigateway', {
      existingApiGatewayObj: apiGateway,
        cloudFrontDistributionProps:{
            defaultBehavior: {
                viewerProtocolPolicy: cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
                allowedMethods: cloudfront.AllowedMethods.ALLOW_ALL,
                cachePolicy: new cloudfront.CachePolicy(this,"cache",{
                    queryStringBehavior: cloudfront.CacheQueryStringBehavior.all(),
                    enableAcceptEncodingGzip: true,
                    enableAcceptEncodingBrotli: true,
                    cookieBehavior: cloudfront.CacheCookieBehavior.all(),
                })
            },
          priceClass: cloudfront.PriceClass.PRICE_CLASS_100,
          domainNames: ["og.hufghani.dev"],
          certificate: certificateManager.Certificate.fromCertificateArn(
              this,
              'certificate',
              ssm.StringParameter.fromStringParameterName(
                  this,
                  'certificateSSM',
                  '/hufghani.dev/Certificate'
              ).stringValue
          ),
        }
    });


    new route53.ARecord(this, 'distributionEntrytoRoute53', {
      target: route53.RecordTarget.fromAlias(
          new aws_route53_targets.CloudFrontTarget(distribution.cloudFrontWebDistribution)
      ),
      zone: route53.HostedZone.fromLookup(this, 'hostzone', {
          domainName: 'hufghani.dev',
      }),
      recordName: "og.hufghani.dev",
    })
  }

}
