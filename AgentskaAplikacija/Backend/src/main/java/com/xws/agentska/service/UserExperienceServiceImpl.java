package com.xws.agentska.service;

import com.xws.agentska.model.Comment;
import com.xws.agentska.model.UserExperience;
import com.xws.agentska.repository.UserExperienceRepository;
import com.xws.agentska.service.interfaces.CompanyService;
import com.xws.agentska.service.interfaces.UserExperienceService;
import com.xws.agentska.service.interfaces.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.sql.Timestamp;
import java.time.Instant;
import java.util.List;

@Service
public class UserExperienceServiceImpl extends CustomGenericService<UserExperience,Long> implements UserExperienceService {
    @Autowired
    private CompanyService companyService;
    @Autowired
    private UsersService usersService;

    @Override
    public void save(UserExperience item) {
        item.setCompany(companyService.findById(item.getCompany().getId()));
        item.setCreatedAt(Timestamp.from(Instant.now()));
        item.setCreator(usersService.findById(item.getCreator().getId()));
        super.save(item);
    }

    @Override
    public List<Comment> findAllCommentsByCompanyId(long id) {
        return ((UserExperienceRepository)repository).findAllCommentsByCompanyId(id);
    }
}
