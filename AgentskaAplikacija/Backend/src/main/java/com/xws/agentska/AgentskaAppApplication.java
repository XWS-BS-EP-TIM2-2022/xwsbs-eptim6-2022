package com.xws.agentska;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@org.springframework.data.jpa.repository.config.EnableJpaRepositories
@SpringBootApplication
public class AgentskaAppApplication {

    public static void main(String[] args) {
        SpringApplication.run(AgentskaAppApplication.class, args);
    }

}
