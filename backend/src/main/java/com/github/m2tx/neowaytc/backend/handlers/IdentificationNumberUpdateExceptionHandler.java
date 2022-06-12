package com.github.m2tx.neowaytc.backend.handlers;

import java.util.Arrays;
import java.util.List;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

import com.github.m2tx.neowaytc.backend.exceptions.IdentificationNumberUpdateException;

@ControllerAdvice
public class IdentificationNumberUpdateExceptionHandler {

    @ExceptionHandler(IdentificationNumberUpdateException.class)
    @ResponseBody
    public ResponseEntity<List<String>> process(final IdentificationNumberUpdateException ex) {
        return new ResponseEntity<List<String>>(Arrays.asList(ex.getMessage()), HttpStatus.BAD_REQUEST);
    }
}
