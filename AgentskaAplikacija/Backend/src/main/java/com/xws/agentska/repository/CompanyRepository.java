package com.xws.agentska.repository;

import com.xws.agentska.model.Company;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

public interface CompanyRepository extends JpaRepository<Company,Long> {
    @Query("select c from Company c where c.owner.id=?1")
    public Company findByOwnerId(long id);
}
