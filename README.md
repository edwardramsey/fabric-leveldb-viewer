# fabric-leveldb-viewer
viewer for hyperledger fabric leveldb

hyperledger fabric peer 的 index 数据【key 在blockfile 的位置】，history key 数据，本地的 channel 信息数据会保存在本地的 leveldb 中，此工具用来辅助读取。

因为 leveldb 只支持单进程读取，peer 启动之后会占用 leveldb，因此使用时需要拷贝 leveldb 目录到新的目录。
