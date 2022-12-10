# TODO

- [ ] 用户允许上传的文件大小限制，避免文件太大导致服务器崩溃
- [ ] 每次上传的文件数量限制，避免一次上传太多，处理不过来
- [ ] 每天上传的文件数量限制，用于避免用户恶意上传文档
- [ ] 根据用户组来进行额度授权
- [ ] 定时清除cache/convert目录下的文件
- [ ] 去除document表中的封面。因为封面这些，都是按约定的方式存储的路径，不需要存储在数据库中
- [ ] 增加一个阅读放大器，类似购物商城那种书标经过，可以放大某一块