##
## $Id: Makefile.am,v 1.2 2001-09-27 12:01:12 chris Exp $
##

AUTOMAKE_OPTIONS = foreign 1.4

PGMS = $(ARCH)/start.o

S_FILES = start.S
S_O_FILES = $(S_FILES:%.S=$(ARCH)/%.o)

OBJS = $(S_O_FILES)

include $(RTEMS_ROOT)/make/custom/@RTEMS_BSP@.cfg
include $(top_srcdir)/../../../../../../automake/lib.am

#
# (OPTIONAL) Add local stuff here using +=
#

$(PROJECT_RELEASE)/lib/start$(LIB_VARIANT).o: $(PGMS)
	$(INSTALL_DATA) $< $@

TMPINSTALL_FILES += $(PROJECT_RELEASE)/lib/start$(LIB_VARIANT).o

all-local: $(ARCH) $(OBJS) $(TMPINSTALL_FILES)

EXTRA_DIST = start.S start.ram start.rom

include $(top_srcdir)/../../../../../../automake/local.am
