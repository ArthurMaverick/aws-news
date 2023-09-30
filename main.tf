resource "aws_iam_role" "rss" {
  name = "rss"
  inline_policy {
    name = "rss-policy"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Action   = ["*"]
          Effect   = "Allow"
          Resource = "*"
        },
      ]
    })
  }
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

data "aws_ssm_parameter" "discord_webhook" {
  name = "/AWS_NEWS/DISCORD/WEBHOOK"
}
data "aws_ssm_parameter" "avatar_url" {
  name = "/AWS_NEWS/DISCORD/AVATAR_URL"
}
data "aws_ssm_parameter" "color" {
  name = "/AWS_NEWS/DISCORD/COLOR"
}

data "external" "rss" {
  program = ["bash", "${path.module}/build.sh"]
}

data "archive_file" "rss" {
  depends_on       = [data.external.rss]
  type             = "zip"
  output_file_mode = "0666"
  source_file      = "${path.module}/main"
  output_path      = "${path.module}/main.zip"
}

resource "aws_lambda_function" "rss" {
  depends_on                     = [data.archive_file.rss]
  filename                       = data.archive_file.rss.output_path
  function_name                  = "rss-mapper"
  description                    = "aws rss"
  handler                        = "main"
  runtime                        = "go1.x"
  architectures                  = ["x86_64"]
  memory_size                    = 128
  timeout                        = 5
  role                           = aws_iam_role.rss.arn
  source_code_hash               = filebase64sha256(data.archive_file.rss.output_path)

  environment {
    variables = {
      DISCORD_WEBHOOK = data.aws_ssm_parameter.discord_webhook.value
      AVATAR_URL      = data.aws_ssm_parameter.avatar_url.value
      COLOR           = data.aws_ssm_parameter.color.value
    }
  }

  timeouts {
    create = "5m"
  }
}

terraform {
  backend "s3" {}
}

