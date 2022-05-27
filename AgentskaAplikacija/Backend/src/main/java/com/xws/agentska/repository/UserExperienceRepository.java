package com.xws.agentska.repository;

import com.xws.agentska.model.Comment;
import com.xws.agentska.model.UserExperience;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;
import java.util.Set;

public interface UserExperienceRepository extends JpaRepository<UserExperience,Long> {
    @Query("select userExp from UserExperience userExp where userExp.company.id=?1 and TYPE(userExp)=Comment")
    public List<Comment> findAllCommentsByCompanyId(long id);
}
