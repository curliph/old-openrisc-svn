/* generic.c -- Definition of generic functions for peripheral to 
 * communicate with host
   
   Copyright (C) 2002 Richard Prescott <rip@step.polymtl.ca
   Copyright (C) 2008 Embecosm Limited

   Contributor Jeremy Bennett <jeremy.bennett@embecosm.com>

   This file is part of Or1ksim, the OpenRISC 1000 Architectural Simulator.

   This program is free software; you can redistribute it and/or modify it
   under the terms of the GNU General Public License as published by the Free
   Software Foundation; either version 3 of the License, or (at your option)
   any later version.

   This program is distributed in the hope that it will be useful, but WITHOUT
   ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
   FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
   more details.

   You should have received a copy of the GNU General Public License along
   with this program.  If not, see <http://www.gnu.org/licenses/>.  */

/* This program is commented throughout in a fashion suitable for processing
   with Doxygen. */


/* Autoconf and/or portability configuration */
#include "config.h"

/* System includes */
#include <stdlib.h>
#include <errno.h>


int
generic_open (void *data)
{
  if (data)
    {
      return 0;
    }
  errno = ENODEV;
  return -1;
}


void
generic_close (void *data)
{
  return;
}


void
generic_free (void *data)
{
  if (data)
    {
      free (data);
    }
}
