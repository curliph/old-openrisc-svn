##
##  $Id: Makefile.am,v 1.2 2001-09-27 12:02:54 chris Exp $
##

AUTOMAKE_OPTIONS = foreign 1.4 no-installman
ACLOCAL_AMFLAGS = -I $(RTEMS_TOPdir)/aclocal

noinst_PROGRAMS = shgen

shgen_SOURCES = sci.h sci.c shgen.c

if HELP2MAN
man_MANS = shgen.1

shgen.1: $(srcdir)/shgen.c
	$(HELP2MAN) -N shgen >$@
endif

include $(top_srcdir)/../../../automake/host.am
