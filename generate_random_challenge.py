import requests
import random
import webbrowser
import os

# Função para buscar uma página específica de desafios
def get_challenges(page_offset, limit=10):
    url = f"https://www.hackerrank.com/rest/contests/master/tracks/algorithms/challenges?offset={page_offset}&limit={limit}&track_login=true"
    headers = {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0",
        "Accept": "application/json"
    }
    
    # Realizar a requisição GET
    response = requests.get(url, headers=headers)
    
    if response.status_code == 200:
        return response.json()
    else:
        print(f"Failed to fetch data. Status code: {response.status_code}")
        return None

# Função para abrir o link no navegador padrão
def open_challenge_in_browser(challenge_slug):
    challenge_url = f"https://www.hackerrank.com/challenges/{challenge_slug}"
    webbrowser.open(challenge_url)  # Abre o navegador padrão

# Função para contar os arquivos em cada pasta
def count_files_in_difficulty_folders():
    base_path = os.path.dirname(os.path.abspath(__file__))  # Diretório onde o script está
    difficulty_folders = ['easy', 'medium', 'hard']
    total_files = 0

    for folder in difficulty_folders:
        folder_path = os.path.join(base_path, folder)
        if os.path.exists(folder_path) and os.path.isdir(folder_path):
            files_in_folder = len([f for f in os.listdir(folder_path) if os.path.isfile(os.path.join(folder_path, f))])
            print(f"Found {files_in_folder} files in '{folder}' folder.")
            total_files += files_in_folder
        else:
            print(f"Folder '{folder}' does not exist.")
    
    return total_files

# Função para verificar se o desafio já foi gerado anteriormente
def is_challenge_already_generated(challenge_name, filename="generated_challenges.txt"):
    if not os.path.exists(filename):
        return False
    with open(filename, 'r') as file:
        generated_challenges = file.read().splitlines()
    return challenge_name in generated_challenges

# Função para salvar o desafio no arquivo .txt
def save_challenge(challenge_name, filename="generated_challenges.txt"):
    with open(filename, 'a') as file:
        file.write(f"{challenge_name}\n")

# Função para buscar e garantir que o desafio não foi gerado antes
def get_unique_challenge(challenges_data):
    while True:
        random_challenge = random.choice(challenges_data)
        challenge_name = random_challenge['name']
        if not is_challenge_already_generated(challenge_name):
            return random_challenge

# Primeiro, buscar o total de desafios com a primeira requisição (offset=0)
initial_data = get_challenges(0, limit=1)  # Limit 1 apenas para pegar o total

if initial_data and "total" in initial_data:
    total_challenges = initial_data["total"]  # Número total de desafios
    print(f"Total Challenges: {total_challenges}")
    
    # Definir número de desafios por página
    challenges_per_page = 10
    
    # Calcular o número total de páginas
    total_pages = total_challenges // challenges_per_page
    
    # Selecionar uma página aleatória
    random_page_offset = random.randint(0, total_pages) * challenges_per_page
    
    # Buscar os desafios da página aleatória
    challenges_data = get_challenges(random_page_offset)
    
    if challenges_data and "models" in challenges_data:
        challenges = challenges_data["models"]
        
        # Garantir que o desafio não foi gerado antes
        random_challenge = get_unique_challenge(challenges)
        
        # Exibir o desafio selecionado
        print(f"Challenge: {random_challenge['name']}")
        print(f"Preview: {random_challenge['preview']}")
        print(f"Difficulty: {random_challenge['difficulty_name']}")
        
        # Abrir o link do desafio no navegador
        open_challenge_in_browser(random_challenge['slug'])
        
        # Sortear a linguagem de programação
        languages = ["Java", "TypeScript", "Python", "Kotlin", "Golang", "Rust"]
        random_language = random.choice(languages)
        print(f"You should solve this challenge using: {random_language}")
        
        # Salvar o nome do desafio no arquivo para evitar repetições
        save_challenge(random_challenge['name'])
        
        # Contar os arquivos nas pastas 'easy', 'medium', e 'hard'
        total_completed = count_files_in_difficulty_folders()
        print(f"Total completed challenges: {total_completed}/{total_challenges}")
    else:
        print("No challenges found.")
else:
    print("Failed to retrieve the total number of challenges.")
