resource "google_service_account" "cloud_run_invoker" {
  account_id   = "${var.project_prefix}-${var.project_env}-cloud-run-invoker"
  display_name = "A service account for cloud run invoke"
}

# google_project_iam_bindingは既存のサービスアカウントのロールを剥奪する恐れがあるため、google_project_iam_memberを使う
# @ref https://zenn.dev/ptiringo/articles/7dd246fcaa73da19d5fb
resource "google_project_iam_member" "cloud_run_invoker_secret_accessor" {
  project = var.project_id
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${google_service_account.cloud_run_invoker.email}"
}
