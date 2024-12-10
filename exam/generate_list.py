import os

# Пути к файлам
questions_files = ["questions_1.txt", "questions_2.txt"]
tasks_dir = "."
output_file = "exam_document.tex"



# Шаблон документа
latex_template = r"""
\documentclass[a4paper,12pt]{article}
\usepackage[utf8]{inputenc}
\usepackage[russian]{babel}
\usepackage{titlesec}
\usepackage{hyperref}
\titleformat{\section}[block]{\normalfont\Large\bfseries}{}{0em}{}
\titleformat{\subsection}[block]{\normalfont\large\bfseries}{}{0em}{}


\title{Вопросы к экзамену по разработке мобильных приложений}
\author{}
\date{}

\begin{document}

\maketitle

\textbf{В билете должно быть два теоретических и один практический вопрос}

\section*{Теоретические вопросы}
%s

\section*{Практические задания}
%s

\section*{Ссылки}

\noindent\url{https://developer.android.com/}

\noindent\url{https://kotlinlang.ru/docs/kotlin-doc.html}

\end{document}
"""

# Считываем теоретические вопросы
theoretical_questions = []
for file in questions_files:
    with open(file, 'r', encoding='utf-8') as f:
        theoretical_questions.append(f.read())

theoretical_section = "\n".join(r"\begin{itemize}" + "\n" +
                                 "".join(f"    \\item {q.strip()}\n" for q in questions.splitlines() if q.strip()) +
                                 r"\end{itemize}" for questions in theoretical_questions)

# Считываем практические задания
practical_tasks = []
for i in range(1, 26):  # Генерируем номера файлов от 1/task.tex до 25/task.tex
    task_file = os.path.join(tasks_dir, f"{i}/task.tex")
    if os.path.exists(task_file):
        with open(task_file, 'r', encoding='utf-8') as f:
            practical_tasks.append(f"\\subsection*{{Задание {i}}}\n" + f.read())

practical_section = "\n".join(practical_tasks)

# Генерируем итоговый документ
with open(output_file, 'w', encoding='utf-8') as f:
    f.write(latex_template % (theoretical_section, practical_section))

print(f"Документ сгенерирован и сохранён в {output_file}")
