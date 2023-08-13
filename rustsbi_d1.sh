picocom -b 115200 -r -l /dev/ttyUSB0  --imap lfcrlf

cargo debug --see

cargo debug --see --kernel ::test

cargo debug --see --kernel z1_new.bin --dt nezha.dts 
ls
cargo debug --see --kernel z1.bin --dt nezha.dts 

cargo debug --spl

cargo debug --see

cargo flash --spl

cargo flash --spl --see

cargo flash --reset --spl --see --boot
