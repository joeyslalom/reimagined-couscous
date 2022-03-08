# reimagined-couscous

go + pubsub + protobuf + grpc


## Service Account to read pubsub

        gcloud iam service-accounts create pubsub-subscriber
        gcloud projects add-iam-policy-binding slalom-2020-293920 --member=serviceAccount:pubsub-subscriber@slalom-2020-293920.iam.gserviceaccount.com --role=roles/pubsub.subscriber
        gcloud iam service-accounts add-iam-policy-binding pubsub-subscriber@slalom-2020-293920.iam.gserviceaccount.com --role=roles/iam.workloadIdentityUser --member=serviceAccount:slalom-2020-293920.svc.id.goog[pubsub/reader]

## GKE Autopilot

* Add CloudShell to authorized networks `echo $(dig +short myip.opendns.com @resolver1.opendns.com)/32`
* Install the chart `helm install rc helm/ -n pubsub --create-namespace`
* Iterate `helm upgrade rc helm/ -n pubsub`

## Cobra
* Install `go install github.com/spf13/cobra/cobra@latest`

        cobra init
        cobra add reader
        cobra add writer

## CloudBuild

* Buildpack https://github.com/GoogleCloudPlatform/buildpacks/blob/main/README.md
* Environment Variable: `GOOGLE_ENTRYPOINT`=`main reader`

## Protobuf + gRPC

* Install and run the `protoc` compiler:

      protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     proto/hello.proto