package com.rhythmo.app.controller;

import org.springframework.boot.web.servlet.error.ErrorController;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

import javax.servlet.http.HttpServletRequest;

@Controller
public class ErrorHandlerController implements ErrorController {

    @RequestMapping("/error")
    public String handleError(HttpServletRequest request) {
        Object status = request.getAttribute("javax.servlet.error.status_code");

        if (status != null) {
            int statusCode = Integer.parseInt(status.toString());
            if (statusCode == 404)
                return "error-404";
            else if (statusCode == 500)
                return "error-500";
        }

        return "error"; // This refers to a view named "error.html" or "error.jsp"
    }

    public String getErrorPath() {
        return "/error";
    }
}
