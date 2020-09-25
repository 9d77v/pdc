# raspi config
dtparam=i2c_arm=on
dtparam=i2s=on
dtparam=spi=on
dtoverlay=i2c-gpio,bus=0,i2c_gpio_delay_us=2,i2c_gpio_sda=23,i2c_gpio_scl=24

# bme280 use bus 1, bh1750 use bus 0