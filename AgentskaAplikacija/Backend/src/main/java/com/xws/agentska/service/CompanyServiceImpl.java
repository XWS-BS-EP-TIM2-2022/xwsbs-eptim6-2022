package com.xws.agentska.service;

import com.xws.agentska.model.*;
import com.xws.agentska.model.enumerations.Status;
import com.xws.agentska.repository.JobPositionRepository;
import com.xws.agentska.service.interfaces.CompanyService;
import com.xws.agentska.service.interfaces.JobPositionService;
import com.xws.agentska.service.interfaces.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.transaction.Transactional;
import java.sql.Timestamp;
import java.time.Instant;

@Service
public class CompanyServiceImpl extends CustomGenericService<Company, Long> implements CompanyService {
    @Autowired
    private UsersService usersService;

    @Autowired
    private JobPositionService jobPositionService;

    @Override
    public void save(Company item) {
        item.setStatus(Status.PENDING);
        item.setOwner(usersService.findById(item.getOwner().getId()));
        super.save(item);
    }

    @Transactional
    @Override
    public void updateCompanyStatus(long companyId, Status newStatus) {
        Company company = this.findById(companyId);
        company.setStatus(newStatus);
        this.update(company);
        usersService.updateUserRole(company.getOwner().getId(), new Role(Role.COMPANY_OWNER_ROLE));
    }

    @Override
    public void addNewJobOffer(long companyId, JobOffer offer) {
        Company company = this.findById(companyId);
        offer.setCreatedAt(Timestamp.from(Instant.now()));
        setJobPositionIfExistsInDb(offer);
        company.addNewJobOffer(offer);
        this.update(company);
    }

    private void setJobPositionIfExistsInDb(JobOffer offer) {
        JobPosition position=jobPositionService.findByName(offer.getPosition().getName());
        if (position!=null) offer.setPosition(position);
    }

    @Override
    public void addNewApiConnection(long companyId, ApiConnection connection) {
        Company company = this.findById(companyId);
        company.addNewApiConnection(connection);
        this.update(company);
    }
}
