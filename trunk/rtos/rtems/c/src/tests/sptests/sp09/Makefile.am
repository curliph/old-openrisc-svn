##
## $Id: Makefile.am,v 1.2 2001-09-27 12:02:31 chris Exp $
##

AUTOMAKE_OPTIONS = foreign 1.4

TEST = sp09

MANAGERS = io message dual_ported_memory event multi_processor region \
    semaphore signal rate_monotonic timer partition

C_FILES = init.c delay.c isr.c screen01.c screen02.c screen03.c screen04.c \
    screen05.c screen06.c screen07.c screen08.c screen09.c screen10.c \
    screen11.c screen12.c screen13.c screen14.c task1.c task2.c task3.c \
    task4.c
C_O_FILES = $(C_FILES:%.c=${ARCH}/%.o)

H_FILES = system.h
noinst_HEADERS = $(H_FILES)

DOCTYPES = scn doc
DOCS = $(DOCTYPES:%=$(TEST).%)

SRCS = $(C_FILES) $(H_FILES)
OBJS = $(C_O_FILES)

PRINT_SRCS = $(DOCS)

PGM = ${ARCH}/$(TEST).exe

include $(RTEMS_ROOT)/make/custom/@RTEMS_BSP@.cfg
include $(RTEMS_ROOT)/make/leaf.cfg
include $(top_srcdir)/sptests.am

${PGM}: $(OBJS) $(LINK_FILES)
	$(make-exe)

all-local: $(ARCH) $(TMPINSTALL_FILES)

EXTRA_DIST = $(C_FILES) $(DOCS)

include $(top_srcdir)/../../../../automake/local.am
