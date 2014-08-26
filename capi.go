// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate mingw32-make -f Makefile.$GOOS

package opencv

/*
#cgo CFLAGS : -I./opencv110/cxcore/include -I./opencv110/cv/include -I./opencv110/cvaux/include  -I./opencv110/ml/include -I./opencv110/highgui/include
#cgo LDFLAGS: -L. -lopencv110

#include <cxcore.h>
#include <cv.h>
*/
import "C"
