obj -m := myHook.o
myHookmodules-objs:=module

KDIR := /lib/modules/2.6.32-279.el6.x86_64/source/
MAKE := make

default:
	$(MAKE) -C $(KDIR) SUBDIRS=$(shell pwd) modules
	
clean:
	$(MAKE) -C $(KDIR) SUBDIRS=$(shell pwd) clean