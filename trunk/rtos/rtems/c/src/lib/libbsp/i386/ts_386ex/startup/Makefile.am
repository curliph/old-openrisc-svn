##
## $Id: Makefile.am,v 1.2 2001-09-27 11:59:50 chris Exp $
##

AUTOMAKE_OPTIONS = foreign 1.4

VPATH = @srcdir@:@srcdir@/../../../shared:@srcdir@/../../shared/comm:@srcdir@/../../shared/irq:@srcdir@/../../shared/io

PGM = $(ARCH)/startup.rel

C_FILES = bspclean.c bsplibc.c bsppost.c bspstart.c bootcard.c main.c sbrk.c \
    irq.c irq_init.c i386-stub.c i386-stub-glue.c uart.c \
    gnatinstallhandler.c
C_O_FILES = $(C_FILES:%.c=$(ARCH)/%.o)

S_FILES = irq_asm.S
S_O_FILES = $(S_FILES:%.S=$(ARCH)/%.o)

startup_rel_OBJECTS = $(C_O_FILES) $(S_O_FILES)

include $(RTEMS_ROOT)/make/custom/@RTEMS_BSP@.cfg
include $(top_srcdir)/../../../../../../automake/lib.am

#
# (OPTIONAL) Add local stuff here using +=
#

## FIXME: This doesn't seem to be useful
## AM_CPPFLAGS += -I$(srcdir)

AM_CPPFLAGS += -DUSE_INIT_FINI

$(PGM): $(startup_rel_OBJECTS)
	$(make-rel)

$(PROJECT_RELEASE)/lib/linkcmds: linkcmds
	$(INSTALL_DATA) $< $@

UNUSED_C_FILES = setvec.c

TMPINSTALL_FILES += $(PROJECT_RELEASE)/lib/linkcmds

all-local: $(ARCH) $(startup_rel_OBJECTS) $(PGM) $(TMPINSTALL_FILES)

.PRECIOUS: $(PGM)

EXTRA_DIST = README bspstart.c linkcmds setvec.c

include $(top_srcdir)/../../../../../../automake/local.am
