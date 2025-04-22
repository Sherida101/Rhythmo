package com.rhythmo.app.controller;

import com.rhythmo.app.model.Habit;
import com.rhythmo.app.controller.HabitController;
import com.rhythmo.app.service.HabitService;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.test.web.servlet.MockMvc;

import java.util.List;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;
import static org.mockito.Mockito.when;

@WebMvcTest(HabitController.class)
class HabitControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private HabitService habitService;

    @Test
    void testGetHabitsByCategory() throws Exception {
        Habit habit = new Habit();
        habit.setId(1L);
        habit.setCategory("Health");

        when(habitService.getHabitsByCategory("Health")).thenReturn(List.of(habit));

        mockMvc.perform(get("/habits/category/Health"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$[0].category").value("Health"));
    }
}
