package com.xws.agentska.service.interfaces;

import com.xws.agentska.model.Comment;
import com.xws.agentska.model.InterviewExperience;
import com.xws.agentska.model.SalaryExperience;
import com.xws.agentska.model.UserExperience;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface UserExperienceService extends Service<UserExperience,Long> {
    public List<Comment> findAllCommentsByCompanyId(long id);
    public List<SalaryExperience> findAllSalariesByCompanyId(long id);
    public List<InterviewExperience> findAllInterviewsByCompanyId(long id);
}
