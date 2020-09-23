package io.chubao.fs.sdk.libsdk;

import com.sun.jna.*;

public interface CFSDriver extends Library {
  public long cfs_new_client();
  public int cfs_set_client(long cid, GoString.ByValue key, GoString.ByValue value);
  public int cfs_start_client(long cid);
  public void cfs_close_client(long cid);
  public int cfs_open(long cid, GoString.ByValue path, int flags, int mode, int uid, int gid, CFSOpenRes.ByReference res);
  public int cfs_flush(long cid, long fd);
  public int cfs_close(long cid, long fd);
  public int cfs_write(long cid, long fd, long offset, byte[] data, int len);
  public int cfs_read(long cid, long fd, long offset, byte[] data, int len);
  public int cfs_mkdirs(long cid, GoString.ByValue path, int mode, int uid, int gid);
  public int cfs_unlink(long cid, GoString.ByValue path);
  public int cfs_rmdir(long cid, GoString.ByValue path, boolean recursive);
  public int cfs_rename(long cid, GoString.ByValue from, GoString.ByValue to);
  public int cfs_getattr(long cid, GoString.ByValue path, SDKStatInfo.ByReference info);
  public int cfs_setattr_by_path(long cid, GoString.ByValue path, SDKStatInfo.ByReference info);
  public int cfs_listattr(long cid, long ino, int num, SDKStatInfo[] info);
  public int cfs_countdir(long cid, GoString.ByValue path, CFSCountDirRes.ByReference res);
  //public int cfs_readdir(long cid, long fd, dirents []C.struct_cfs_dirent, count int) (n int) {

}