package com.xws.agentska.service;

import com.xws.agentska.model.Company;
import com.xws.agentska.model.Role;
import com.xws.agentska.model.enumerations.Status;
import com.xws.agentska.service.interfaces.CompanyService;
import com.xws.agentska.service.interfaces.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.transaction.Transactional;

@Service
public class CompanyServiceImpl extends CustomGenericService<Company,Long> implements CompanyService {
    @Autowired
    private UsersService usersService;
    @Override
    public void save(Company item) {
        item.setStatus(Status.PENDING);
        item.setOwner(usersService.findById(item.getOwner().getId()));
        super.save(item);
    }
    @Transactional
    @Override
    public void updateCompanyStatus(long companyId, Status newStatus) {
        Company company= this.findById(companyId);
        company.setStatus(newStatus);
        this.update(company);
        usersService.updateUserRole(company.getOwner().getId(), new Role(Role.COMPANY_OWNER_ROLE));
    }
}
