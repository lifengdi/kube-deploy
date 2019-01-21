echo "$f"
echo "$kubeconfs"
echo "$imagePullSecrets"
echo "$log"

kube-deploy -kubeconfs "$kubeconfs" -f "$f" -imagePullSecrets "$imagePullSecrets" -log "$log"