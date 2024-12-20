import os
import random

def load_questions(file_path):
    """Load questions from a file, one per line."""
    with open(file_path, "r", encoding="utf-8") as f:
        return f.readlines()

def load_tasks(task_directory):
    """Load practical tasks from a directory."""
    tasks = []
    for i in range(1, 26):
        task_path = os.path.join(task_directory, f"{i}/task.tex")
        with open(task_path, "r", encoding="utf-8") as f:
            tasks.append(f.read().strip())
    return tasks

def generate_tickets(tickets_template_path, ticket_template_path, questions_1, questions_2, tasks, output_path):
    """Generate LaTeX file with 25 exam tickets."""
    # Shuffle questions
    shuffled_questions_1 = random.sample(questions_1, len(questions_1))
    shuffled_questions_2 = random.sample(questions_2, len(questions_2))

    # Load templates
    with open(tickets_template_path, "r", encoding="utf-8") as f:
        tickets_template = f.read()

    with open(ticket_template_path, "r", encoding="utf-8") as f:
        ticket_template = f.read()

    # Generate tickets
    tickets_content = []
    for i in range(25):
        q1 = shuffled_questions_1[i].strip()
        q2 = shuffled_questions_2[i].strip()
        q3 = tasks[i].replace("\\begin{verbatim}","").replace("\\end{verbatim}""","").replace("&","\\&").replace("_","\\_").replace(" {"," \\{").replace(" }"," \\}")
        ticket_content = ticket_template.replace("@num@", str(i+1)).replace("@q1@", q1).replace("@q2@", q2).replace("@q3@", q3)
        tickets_content.append(ticket_content)

    # Insert tickets into the main template
    full_content = tickets_template.replace("@content@", "\n\n".join(tickets_content))

    # Write output file
    with open(output_path, "w", encoding="utf-8") as f:
        f.write(full_content)

# Example usage
if __name__ == "__main__":
    questions_1_path = "questions_1.txt"
    questions_2_path = "questions_2.txt"
    tasks_directory = ""
    tickets_template_path = "tickets.tex"
    ticket_template_path = "ticket.tex"
    output_path = "generated_tickets.tex"

    # Load data
    questions_1 = load_questions(questions_1_path)
    questions_2 = load_questions(questions_2_path)
    tasks = load_tasks(tasks_directory)

    # Generate tickets
    generate_tickets(tickets_template_path, ticket_template_path, questions_1, questions_2, tasks, output_path)

    print(f"Tickets generated in {output_path}")
