作为即将毕业的一名运维从业人员，我对云原生发展方向很感兴趣，特别看好k8s。

经过了15周的学习，对k8s有了一定的了解：

- 在基础概念方面，分清了k8s基础原生组件的一些作用。
  - api-server，作为唯一和etcd通信的入口，对请求认证、授权、准入（mutating，修改请求的参数），管理原生的所有对象。
  - etcd，集群数据的存储后端（持久化分布式键值存储，奇数节点），集群的高可用约等于etcd高可用，支持watch机制，例如 kube-controller-manager 和 kube-scheduler 都通过watch机制（通过watch api-server）来更新一些内部状态。
  - kube-scheduler，根据资源需求、亲和性、反亲和性、调度器策略等来进行节点选择。scheduler基本是一个非常稳定的组件。
  - kube-controller-manager，多个控制器的组合，例如StatefulSet/Replication/CronJob/DaemonSet。kcm通过apiserver watch到后，就开始干活了。还提供了垃圾回收、节点监控等功能。
  - kubelet，定期向api-server报告pod/节点状态，监控并负责pod的运行，通过自动重启或终止，确保pod按照期望运行。拥有当前节点的最大控制权。
  - kube-proxy，负责维护节点的网络规则和转发网络请求（iptables，ipvs），从而确保容器之间可以互相通信，并且向外部网络提供服务。和网络插件作用范围不同，例如Calico 更多地关注于容器网络的实现和管理，而 kube-proxy 主要作用于节点级别的网络转发和负载均衡服务。

- 在掌握核心技能方面，我感觉自己有稳步进步。
  - golang的基本语法，docker镜像的封装。
  - 创建Deployment、ConfigMap、Service，并考虑到了容器健康（避免/bin/sh启动容器，优雅退出可使用tini），自动扩缩容，亲和性等方面。
  - 当流量进来时，使用istio进行分发和灰度，通过PeerAuthentication/PeerAuthentication（前者用于服务之间的身份验证，比如强制mtls，确保服务之间建立的通信是安全的；后者用于请求身份验证，在确保请求是来自合法身份），AuthorizationPolicy对象进行认证，鉴权。
  - 学会了使用helm部署应用（修改values配置文件比较麻烦），能看懂一些基本的模版文件。
  - 学会了使用loki查看日志，配置grafana面板。

- 同时也了解了一些生产环境中遇到的坑和建议。
  - 权限最小化