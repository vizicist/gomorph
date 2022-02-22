#include <stdio.h>
#include <sensel.h>

SenselDeviceList sensellist;

int main()
{
	senselGetDeviceList(&sensellist);
	int ndevices = sensellist.num_devices;
	printf("ndevices = %d\n",ndevices);
	int n;
	for ( n=0; n<ndevices; n++ ) {
		printf("n=%d serial_num=%s\n",n,sensellist.devices[n].serial_num);
		unsigned char idx = sensellist.devices[n].idx;
		SENSEL_HANDLE h;
		printf("Trying to open device with idx = %d\n",idx);
		SenselStatus status = senselOpenDeviceByID(&h,idx);
		printf("status = %d  handle = %lld\n",status,(long long)h);
	}
	return(0);
}
