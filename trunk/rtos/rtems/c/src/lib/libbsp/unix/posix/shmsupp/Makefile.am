##
## $Id: Makefile.am,v 1.2 2001-09-27 12:01:15 chris Exp $
##

AUTOMAKE_OPTIONS = foreign 1.4

PGM = $(ARCH)/shmsupp.rel

C_FILES = addrconv.c getcfg.c cause_intr.c lock.c mpisr.c
C_O_FILES = $(C_FILES:%.c=$(ARCH)/%.o)

OBJS = $(C_O_FILES)

include $(RTEMS_ROOT)/make/custom/@RTEMS_BSP@.cfg
include $(top_srcdir)/../../../../../../automake/lib.am

#
# (OPTIONAL) Add local stuff here using +=
#

$(PGM): $(OBJS)
	$(make-rel)

if HAS_MP
all-local: $(ARCH) $(OBJS) $(PGM)
endif

# the .rel file built here will be put into libbsp.a by ../wrapup/Makefile

.PRECIOUS: $(PGM)

EXTRA_DIST = README addrconv.c cause_intr.c getcfg.c lock.c mpisr.c

include $(top_srcdir)/../../../../../../automake/local.am
