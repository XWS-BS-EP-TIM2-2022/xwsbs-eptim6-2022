package com.xws.agentska.controller;

import com.xws.agentska.model.JobPosition;
import com.xws.agentska.service.interfaces.JobPositionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class JobPositionController {
    @Autowired
    private JobPositionService service;
    @GetMapping("/api/job-positions")
    public ResponseEntity<?> findAll(){
        return ResponseEntity.ok(service.findAll());
    }
}
