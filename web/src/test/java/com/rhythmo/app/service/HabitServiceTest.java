package com.rhythmo.app.service;

import com.rhythmo.app.model.Habit;
import com.rhythmo.app.repository.HabitRepository;
import com.rhythmo.app.service.HabitService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

import java.util.Optional;

import static org.mockito.Mockito.*;
import static org.junit.jupiter.api.Assertions.*;

class HabitServiceTest {

    private HabitService habitService;
    private HabitRepository habitRepository;

    @BeforeEach
    void setUp() {
        habitRepository = mock(HabitRepository.class);
        habitService = new HabitService(habitRepository);
    }

    @Test
    void testGetHabitByIdReturnsHabit() {
        Habit mockHabit = new Habit();
        mockHabit.setId(1L);
        mockHabit.setName("Test Habit");

        when(habitRepository.findById(1L)).thenReturn(Optional.of(mockHabit));

        Habit result = habitService.getHabitById(1L);
        assertNotNull(result);
        assertEquals("Test Habit", result.getName());
    }

    @Test
    void testUpdateHabitFrequency() {
        Habit mockHabit = new Habit();
        mockHabit.setId(2L);
        mockHabit.setFrequency("Daily");

        when(habitRepository.findById(2L)).thenReturn(Optional.of(mockHabit));
        when(habitRepository.save(any(Habit.class))).thenReturn(mockHabit);

        habitService.updateHabitFrequency(2L, "Weekly");
        assertEquals("Weekly", mockHabit.getFrequency());
    }
}
