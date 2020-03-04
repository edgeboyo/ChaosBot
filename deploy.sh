gcloud app deploy

gcloud app versions delete $(gcloud app versions list | grep "0\.00" | cut -d " " -f 3) 
