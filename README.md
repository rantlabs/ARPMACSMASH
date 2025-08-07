# ARPMACSMASH
#### ARPMACSMASH Locate all of your network connected devices.
#### ARPMACSMASH is a simple and fast way to join an arp table and a mac table together. Both the arp and mac tables can be any format from any vendor. ARPMACSMASH also takes any mac address format including different formats on the two files you are joining. It takes about a second to join a enterprise campus sized arp and mac table. ARPMACSMASH creates a seperate entry for each MAC table entry which will create multiple ARPMAC entries if you are joining an arp table to a network with multiple layers of L2 switches. You will need to determine the final connection point of a device.

#### USE CASE - Different layer three and layer two vendors. You may have a layer two switch network that connects to layer three firewalls. ARPMACSMASH will connect the two tables to help you inventory and located devices.

#### For command line options type arpmacsmash or arpmacsmash -help
````
arpmacsmash
Both -arpfile and -macfile must be specified
Usage of ./arpmacsmash:
  -arpfile string
    	Path to the ARP table file
  -macfile string
    	Path to the MAC address file
  -output string
    	Path to the output file (optional)
````
#### Executables binaries available for download. The notes.txt file contains source compiling instructions. You can also rename the binary.
````
arpmacsmash_linux_32
arpmacsmash_linux_64
arpmacsmash_mac
arpmacsmash_mac_arm64
arpmacsmash_rpi_arm64
arpmacsmash_rpi_armv6
arpmacsmash_rpi_armv7
arpmacsmash_windows_32.exe
arpmacsmash_windows_64.exe
````
#### Usage example
````
NOTE: The example shows device 192.168.20.11 is connected to a daisy chain of layer 2 switches. The actual device is connected to switch4
arpfile.txt contents:
switchY |show ip arp| Internet  192.168.20.11            64f6.9d29.d538
switchZ |show ip arp| Internet  192.168.10.10              5ce2.86f5.eb67

macfile.txt contents
switchY |show mac address-table|  140    64f6.9d29.d538    DYNAMIC     Port1
switch2 |show mac address-table|  140    64f6.9d29.d538    DYNAMIC     Port2
switch3 |show mac address-table|  140    64f6.9d29.d538    DYNAMIC     Port3
switch4 |show mac address-table|  140    64f6.9d29.d538    DYNAMIC     Port4
switchZ |show mac address-table|  145    5ce2.86f5.eb67    STATIC      Gi1/0/9

arpmacsmash -arpfile arpfile.txt -macfile macfile.txt            
Loaded 2 MAC addresses from ARP file into lookup table
switchY |show ip arp| Internet  192.168.20.11            64f6.9d29.d538 switchY |show mac address-table|  140    64f6.9d29.d538    DYNAMIC     Port1
switchY |show ip arp| Internet  192.168.20.11            64f6.9d29.d538 switch2 |show mac address-table|  140    64f6.9d29.d538    DYNAMIC     Port2
switchY |show ip arp| Internet  192.168.20.11            64f6.9d29.d538 switch3 |show mac address-table|  140    64f6.9d29.d538    DYNAMIC     Port3
switchY |show ip arp| Internet  192.168.20.11            64f6.9d29.d538 switch4 |show mac address-table|  140    64f6.9d29.d538    DYNAMIC     Port4
switchZ |show ip arp| Internet  192.168.10.10            5ce2.86f5.eb67 switchZ |show mac address-table|  145    5ce2.86f5.eb67    STATIC      Gi1/0/9
````
