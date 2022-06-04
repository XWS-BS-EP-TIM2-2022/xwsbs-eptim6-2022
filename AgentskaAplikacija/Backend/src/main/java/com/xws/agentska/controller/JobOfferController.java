package com.xws.agentska.controller;

import com.xws.agentska.service.interfaces.JobOfferService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class JobOfferController {
    @Autowired
    private JobOfferService service;

    @GetMapping("/api/job-offers")
    public ResponseEntity<?> findAll() {
        return ResponseEntity.ok().body(service.findAll());
    }
}
