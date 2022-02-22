#include <stdio.h>
#include <sensel.h>

SenselDeviceList sensellist;

int main()
{

	printf("Hello World\n");
	senselGetDeviceList(&sensellist);
	int ndevices = sensellist.num_devices;
	printf("%d\n",ndevices);
	int n;
	for ( n=0; n<ndevices; n++ ) {
		printf("opening %s\n",sensellist.devices[n].serial_num);
		unsigned char idx = sensellist.devices[n].idx;
		SENSEL_HANDLE h;
		SenselStatus status = senselOpenDeviceByID(&h,idx);
		printf("status = %d\n",status);
	}
	return(0);
}
