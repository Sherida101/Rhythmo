/*
 Rhythmo — Build better habits with rhythm
 https://github.com/Sherida101/Rhythmo

 See LICENSE for copyright details.
*/

package com.rhythmo.app;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan(basePackages = "com.rhythmo.app.controller")
public class RhythmoApplication {

	public static void main(String[] args) {
		SpringApplication.run(RhythmoApplication.class, args);
	}

}
