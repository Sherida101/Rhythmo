import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.http.ResponseEntity;
import java.util.List;
import com.rhythmo.app.model.Habit;

@RestController
@RequestMapping("/api")
public class HabitController {

    @Autowired
    private HabitService habitService;

    @GetMapping("/habits")
    public ResponseEntity<List<Habit>> getHabits() {
        return ResponseEntity.ok(habitService.getAllHabits());
    }
}
