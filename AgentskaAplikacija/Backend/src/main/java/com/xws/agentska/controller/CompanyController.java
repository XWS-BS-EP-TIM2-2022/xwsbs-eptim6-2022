package com.xws.agentska.controller;

import com.xws.agentska.model.Company;
import com.xws.agentska.model.User;
import com.xws.agentska.security.TokenBasedAuthentication;
import com.xws.agentska.service.interfaces.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import static com.xws.agentska.controller.UserController.getLoggedinUserId;

@RestController
public class CompanyController {
    @Autowired
    private CompanyService service;

    @PostMapping("/api/companies")
    public ResponseEntity<?> createNewCompanyRequest(@RequestBody Company company) {
        company.setOwner(new User(getLoggedinUserId()));
        service.save(company);
        return ResponseEntity.ok().build();
    }
    @PutMapping("/api/companies")
    public ResponseEntity<?> createResponseToNewCompanyRequest(@RequestBody Company company) {
        service.updateCompanyStatus(company.getId(),company.getStatus());
        return ResponseEntity.ok().build();
    }

}
