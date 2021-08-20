syscall::setsockopt:entry
/execname=="client" || execname=="server"/
{
    printf("%s::%s(%d,0x%X,0x%X,[%d],%d)",execname,probefunc,arg0,arg1,arg2,*(uint8_t *)copyin(arg3,arg4),arg4);
    ustack();
}
