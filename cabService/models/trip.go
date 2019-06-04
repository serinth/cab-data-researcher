package models

import (
	"time"
)

type Trip struct {
	Medallion      string `xorm:"TEXT"`
	PickupDatetime time.Time
}

//
//
//CREATE TABLE `cab_trip_data` (
//`medallion` text,
//`hack_license` text,
//`vendor_id` text,
//`rate_code` int(11) DEFAULT NULL,
//`store_and_fwd_flag` text,
//`pickup_datetime` datetime DEFAULT NULL,
//`dropoff_datetime` datetime DEFAULT NULL,
//`passenger_count` int(11) DEFAULT NULL,
//`trip_time_in_secs` int(11) DEFAULT NULL,
//`trip_distance` double DEFAULT NULL,
//`pickup_longitude` double DEFAULT NULL,
//`pickup_latitude` double DEFAULT NULL,
//`dropoff_longitude` double DEFAULT NULL,
//`dropoff_latitude` double DEFAULT NULL
//) ENGINE=InnoDB DEFAULT CHARSET=latin1;
