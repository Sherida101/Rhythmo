/*
 Rhythmo — Build better habits with rhythm
 https://github.com/Sherida101/Rhythmo

 See LICENSE for copyright details.
*/

package com.rhythmo.app.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class HomeController {
    @GetMapping("/")
    public String home(Model model) {
        model.addAttribute("message", "Welcome to Rhythmo, a habit tracker web platform!");
        return "index"; // maps to templates/index.html
    }
}
