package morph

// #cgo LDFLAGS: -L. "/Users/tjt/Downloads/SenselStaticMacLib/m1/LibSenselStatic.a"
/*
#include <stdlib.h>
#include <stdio.h>
#include "../SenselLib/include/sensel.h"

// THIS STRUCTURE NEEDS TO MATCH THE C VERSION
// EVENTUALLY I SHOULD GET RID OF THIS
typedef struct goSenselSensorInfo
{
    unsigned char   max_contacts;       // Maximum number of contacts the sensor supports
    unsigned short  num_rows;           // Total number of rows
    unsigned short  num_cols;           // Total number of columns
    float           width;              // Width of the sensor in millimeters
    float           height;             // Height of the sensor in millimeters
} goSenselSensorInfo;

// THIS STRUCTURE DOES NOT NEED TO MATCH THE C VERSION
typedef struct goSenselFirmwareInfo
{
    unsigned char  fw_protocol_version; // Sensel communication protocol supported by the device
    unsigned char  fw_version_major;    // Major version of the firmware
    unsigned char  fw_version_minor;    // Minor version of the firmware
    unsigned short fw_version_build;    // ??
    unsigned char  fw_version_release;  // ??
    unsigned short device_id;           // Sensel device type
    unsigned char  device_revision;     // Device revision
} goSenselFirmwareInfo;

// THIS STRUCTURE DOES NOT NEED TO MATCH THE C VERSION
typedef struct goSenselFrameData
{
    unsigned char   content_bit_mask;  // Data contents of the frame
    int             lost_frame_count;  // Number of frames dropped
    unsigned char   n_contacts;        // Number of contacts
} goSenselFrameData;

// THIS STRUCTURE DOES NOT NEED TO MATCH THE C VERSION
typedef struct goSenselContact
{
    // unsigned char        content_bit_mask;   // Mask of what contact data is valid
    unsigned char        id;                 // Contact id
    unsigned int         state;              // Contact state (enum SenselContactState)
    float                x_pos;              // X position in mm
    float                y_pos;              // Y position in mm
    float                total_force;        // Total contact force in grams
    float                area;               // Area in sensor elements
} goSenselContact;

typedef struct OneMorph {
	// void*			handle;
	SENSEL_HANDLE		handle;
	SenselFrameData *frame;
} OneMorph;

// The order in this list is the idx value
OneMorph Morphs[SENSEL_MAX_DEVICES];

int senselLoaded = 0;
SenselDeviceList sensellist;

void
SenselInit()
{
	if ( ! senselLoaded ) {
		senselGetDeviceList(&sensellist);
		printf("SenselInit: num_devices = %d\n", sensellist.num_devices);
		senselLoaded = 1;
	}
}

int
SenselNumDevices()
{
	SenselInit();
	return sensellist.num_devices;
}

char *
SenselDeviceSerialNum(unsigned char idx) {
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return "InvalidDeviceIndex";
	}
	return (char *)(sensellist.devices[idx].serial_num);
}

// int
// SenselOpenDeviceBySerialNum(void** handle, char* serial_num)
// {
// 	return senselOpenDeviceBySerialNum(handle,serial_num);
// }

int
SenselOpenDeviceByID(unsigned char idx)
{
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return SENSEL_ERROR;
	}
	SENSEL_HANDLE h;
	SenselStatus status = senselOpenDeviceByID(&h,idx);
	Morphs[idx].handle = h;
	return status;
}

int
SenselGetSensorInfo(unsigned char idx, goSenselSensorInfo *info)
{
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return SENSEL_ERROR;
	}
	SENSEL_HANDLE handle = Morphs[idx].handle;
	SenselSensorInfo *senselinfo = (SenselSensorInfo*)info;
	return senselGetSensorInfo(handle,senselinfo);
}

int
SenselGetFirmwareInfo(unsigned char idx, goSenselFirmwareInfo *goinfo)
{
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return SENSEL_ERROR;
	}
	SENSEL_HANDLE handle = Morphs[idx].handle;
	SenselFirmwareInfo info;
	SenselStatus s = senselGetFirmwareInfo(handle,&info);
	if ( s == SENSEL_OK ) {
		goinfo->fw_protocol_version = info.fw_protocol_version;
		goinfo->fw_version_major = info.fw_version_major;
		goinfo->fw_version_minor = info.fw_version_minor;
		goinfo->fw_version_build = info.fw_version_build;
		goinfo->fw_version_release = info.fw_version_release;
		goinfo->device_id = info.device_id;
		goinfo->device_revision = info.device_revision;
	}
	return s;
}

int
SenselSetupAndStart(unsigned char idx)
{
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return -1;
	}
	SENSEL_HANDLE handle = Morphs[idx].handle;

	unsigned char v[1] = {255};
	SenselStatus s;
	s = senselWriteReg(handle,0xD0,1,v);  // This disables timeouts
	if ( s != SENSEL_OK ) {
		return -2;
	}

	s = senselSetFrameContent(handle, FRAME_CONTENT_CONTACTS_MASK);
	if ( s != SENSEL_OK ) {
		return -3;
	}

	s = senselAllocateFrameData(handle, &(Morphs[idx].frame));
	if ( s != SENSEL_OK ) {
		return -4;
	}

	s = senselStartScanning(handle);
	if ( s != SENSEL_OK ) {
		return -5;
	}

	for (int led = 0; led < 16; led++) {
		s = senselSetLEDBrightness(handle, led, 0); //turn off LED
		if ( s != SENSEL_OK ) {
			return -6;
		}
	}
	return SENSEL_OK;
}

int
SenselReadSensor(unsigned char idx)
{
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return SENSEL_ERROR;
	}
	SENSEL_HANDLE handle = Morphs[idx].handle;
	SenselStatus s = senselReadSensor(handle);
	return s;
}

int
SenselGetNumAvailableFrames(unsigned char idx)
{
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return SENSEL_ERROR;
	}
	SENSEL_HANDLE handle = Morphs[idx].handle;
	unsigned int nframes;
	SenselStatus s = senselGetNumAvailableFrames(handle,&nframes);
	if ( s != SENSEL_OK ) {
		return -1;
	}
	return nframes;
}

int
SenselGetFrame(unsigned char idx, goSenselFrameData *goFrame)
{
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return SENSEL_ERROR;
	}
	SENSEL_HANDLE handle = Morphs[idx].handle;
	SenselStatus s = senselGetFrame(handle,Morphs[idx].frame);
	SenselFrameData *f = Morphs[idx].frame;
	goFrame->n_contacts = f->n_contacts;
    goFrame->lost_frame_count = f->lost_frame_count;
    goFrame->content_bit_mask = f->content_bit_mask;
	return s;
}

int
SenselGetContact(unsigned char idx, unsigned char contactNum, goSenselContact *goContact)
{
	if ( idx < 0 || idx >= SENSEL_MAX_DEVICES ) {
		return SENSEL_ERROR;
	}
	SenselFrameData *f = Morphs[idx].frame;
	if ( contactNum >= f->n_contacts ) {
		return SENSEL_ERROR;
	}
	goContact->id = f->contacts[contactNum].id;
	goContact->x_pos = f->contacts[contactNum].x_pos;
	goContact->y_pos = f->contacts[contactNum].y_pos;
	goContact->state = f->contacts[contactNum].state;
	goContact->total_force = f->contacts[contactNum].total_force;
	goContact->area = f->contacts[contactNum].area;
	return SENSEL_OK;
}
*/
import "C"

import (
	"fmt"
	"log"
)

func (m OneMorph) readFrames(callback CursorDeviceCallbackFunc, forceFactor float32) (err error) {
	status := C.SenselReadSensor(C.uchar(m.Idx))
	if status != C.SENSEL_OK {
		return fmt.Errorf("SenselReadSensor for idx=%d returned %d", m.Idx, status)
	}
	numFrames := C.SenselGetNumAvailableFrames(C.uchar(m.Idx))
	if numFrames <= 0 {
		return nil
	}
	// log.Printf("Morph: FRAMES ARE AVAILABLE!! idx=%d numFrames=%d\n", m.Idx, numFrames)
	nf := int(numFrames)
	for n := 0; n < nf; n++ {
		var frame C.struct_goSenselFrameData
		status := C.SenselGetFrame(C.uchar(m.Idx), &frame)
		if status != C.SENSEL_OK {
			return fmt.Errorf("SenselGetFrame of idx=%d returned %d\n", m.Idx, status)
		}
		nc := int(frame.n_contacts)
		for n := 0; n < nc; n++ {
			var contact C.struct_goSenselContact
			status = C.SenselGetContact(C.uchar(m.Idx), C.uchar(n), &contact)
			if status != C.SENSEL_OK {
				return fmt.Errorf("SenselGetContact of morph_idx=%d n=%d returned %d\n", m.Idx, n, status)
			}
			xNorm := float32(contact.x_pos) / m.Width
			yNorm := float32(contact.y_pos) / m.Height
			zNorm := float32(contact.total_force) / MaxForce
			zNorm *= forceFactor
			area := float32(contact.area)
			var ddu string
			switch contact.state {
			case CursorDown:
				ddu = "down"
			case CursorDrag:
				ddu = "drag"
			case CursorUp:
				ddu = "up"
			default:
				return fmt.Errorf("Morph: Invalid value for contact.state - %d\n", contact.state)
				continue
			}

			cid := fmt.Sprintf("%d", contact.id)

			if DebugMorph {
				log.Printf("DebugMorph: serial=%s contact_id=%d morph_idx=%d n=%d state=%d xNorm=%f yNorm=%f zNorm=%f\n",
					m.SerialNum, contact.id, m.Idx, n, contact.state, xNorm, yNorm, zNorm)
			}

			// make the coordinate space match OpenGL and Freeframe
			yNorm = 1.0 - yNorm

			// Make sure we don't send anyting out of bounds
			if yNorm < 0.0 {
				yNorm = 0.0
			} else if yNorm > 1.0 {
				yNorm = 1.0
			}
			if xNorm < 0.0 {
				xNorm = 0.0
			} else if xNorm > 1.0 {
				xNorm = 1.0
			}

			ev := CursorDeviceEvent{
				CID:       cid,
				Timestamp: 0,
				Ddu:       ddu,
				X:         xNorm,
				Y:         yNorm,
				Z:         zNorm,
				Area:      area,
			}
			callback(ev)
		}
	}
	return nil
}

// Initialize xxx
func initialize(serial string) ([]OneMorph, error) {

	C.SenselInit()

	numdevices := int(C.SenselNumDevices())

	morphs := make([]OneMorph, 0)

	for idx := uint8(0); idx < uint8(numdevices); idx++ {

		thisSerial := C.GoString(C.SenselDeviceSerialNum(C.uchar(idx)))

		if serial != "*" && thisSerial != serial {
			continue
		}

		m := OneMorph{}
		m.Idx = idx
		m.SerialNum = thisSerial

		status := C.SenselOpenDeviceByID(C.uchar(idx))
		if status != C.SENSEL_OK {
			return nil, fmt.Errorf("SenselOpenDeviceByID of idx=%d returned %d", idx, status)
		}

		var sensorinfo C.struct_goSenselSensorInfo
		status = C.SenselGetSensorInfo(C.uchar(idx), &sensorinfo)
		if status != C.SENSEL_OK {
			return nil, fmt.Errorf("SenselGetSensorInfo of idx=%d returned %d", idx, status)
		}

		var firmwareinfo C.struct_goSenselFirmwareInfo
		status = C.SenselGetFirmwareInfo(C.uchar(idx), &firmwareinfo)
		if status != C.SENSEL_OK {
			return nil, fmt.Errorf("SenselGetFirmwareInfo of %s returned %d", m.SerialNum, status)
		}

		status = C.SenselSetupAndStart(C.uchar(m.Idx))
		if status != C.SENSEL_OK {
			return nil, fmt.Errorf("SenselSetupAndStart of %s returned %d", m.SerialNum, status)
		}

		m.Opened = true
		m.Width = float32(sensorinfo.width)
		m.Height = float32(sensorinfo.height)
		m.FwVersionMajor = uint8(firmwareinfo.fw_version_major)
		m.FwVersionMinor = uint8(firmwareinfo.fw_version_minor)
		m.FwVersionBuild = uint8(firmwareinfo.fw_version_build)
		m.FwVersionRelease = uint8(firmwareinfo.fw_version_release)
		m.DeviceID = int(firmwareinfo.device_id)

		morphs = append(morphs, m)

	}
	if len(morphs) == 0 {
		return nil, fmt.Errorf("could not open any Morph matching serial=%s", serial)
	}
	return morphs, nil
}
