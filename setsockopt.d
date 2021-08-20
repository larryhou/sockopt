syscall::setsockopt:entry
/execname=="client" || execname=="server"/
{
    self->arg0 = arg0;
    self->arg1 = arg1;
    self->arg2 = arg2;
    self->arg3 = *(char *)copyin(arg3,arg4);
    self->arg4 = arg4;
}

syscall::setsockopt:return
/execname=="client" || execname=="server"/
{
    printf("%s::%s(%d,0x%X,0x%X,[%d],%d) = %d",execname,probefunc,self->arg0,self->arg1,self->arg2,self->arg3,self->arg4,arg1);
    ustack();
}
