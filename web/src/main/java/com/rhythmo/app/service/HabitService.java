import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import java.util.List;
import com.rhythmo.app.model.Habit;
import com.rhythmo.app.repository.HabitRepository;
import java.time.LocalDate;
import java.util.HashMap;
import java.util.Map;
import java.util.function.BiConsumer;
import java.util.function.Function;

@Service
public class HabitService {

    @Autowired
    private final HabitRepository habitRepository;

    public HabitService(HabitRepository habitRepository) {
        this.habitRepository = habitRepository;

        // Initialise updaters and getters
        initialiseGetters();
        initialiseUpdaters();
    }

    private final Map<String, BiConsumer<Habit, Object>> updaters = new HashMap<>();
    private final Map<String, Function<Object, List<Habit>>> getters = new HashMap<>();

    private void initialiseGetters() {
        getters.put("userId", value -> habitRepository.findByUserId((Long) value));
        getters.put("category", value -> habitRepository.findByCategory((String) value));
        getters.put("frequency", value -> habitRepository.findByFrequency((String) value));
        getters.put("status", value -> habitRepository.findByStatus((String) value));
        getters.put("tag", value -> habitRepository.findByTag((String) value));
        getters.put("priority", value -> habitRepository.findByPriority((String) value));
    }

    private void initialiseUpdaters() {
        updaters.put("status", (habit, value) -> habit.setStatus((String) value));
        updaters.put("frequency", (habit, value) -> habit.setFrequency((String) value));
        updaters.put("category", (habit, value) -> habit.setCategory((String) value));
    }

    // Create
    public Habit createHabit(Habit habit) {
        return habitRepository.save(habit);
    }

    // Delete
    public void deleteHabit(Long id) {
        habitRepository.deleteById(id);
    }

    // Get
    public List<Habit> getAllHabits() {
        return habitRepository.findAll();
    }

    public Habit getHabitById(Long id) {
        return habitRepository.findById(id).orElse(null);
    }

    public List<Habit> getHabitsByField(String fieldName, Object value) {
        if (getters.containsKey(fieldName)) {
            return getters.get(fieldName).apply(value);
        }
        return List.of(); // return empty list or throw exception
    }

    // Update
    public Habit updateHabit(Long id, Habit habit) {
        if (habitRepository.existsById(id)) {
            habit.setId(id);
            return habitRepository.save(habit);
        }
        return null;
    }

    // e.g. habitService.updateHabitField(1L, "status", "active");
    public void updateHabitField(Long id, String fieldName, Object value) {
        Habit habit = getHabitById(id);
        if (habit != null && updaters.containsKey(fieldName)) {
            updaters.get(fieldName).accept(habit, value);
            habitRepository.save(habit);
        }
    }
}
