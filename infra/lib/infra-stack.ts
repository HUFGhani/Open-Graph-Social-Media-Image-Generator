import * as cdk from 'aws-cdk-lib';
import {
  aws_apigateway as apigw, aws_certificatemanager as certificatemanager, aws_cloudfront,
  aws_lambda as lambda, aws_route53, aws_route53_targets,
  aws_s3,
  aws_s3_assets as asserts, aws_ssm as ssm,
  Duration,
  RemovalPolicy
} from 'aws-cdk-lib';
import {Construct} from 'constructs';
import * as path from "path";
import {ContentHandling} from "aws-cdk-lib/aws-apigateway";
import {CloudFrontToApiGateway} from "@aws-solutions-constructs/aws-cloudfront-apigateway";


export class InfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

  const lambdaAsset = new asserts.Asset(this, "HelloGoServerLambdaFnZip",{
    path: path.join(__dirname,"../../functions")
  })

    const opengraphAsset = new aws_s3.Bucket(this,'opengraphAsset',{
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
          priceClass: aws_cloudfront.PriceClass.PRICE_CLASS_100,
        }
    });



  }
}
