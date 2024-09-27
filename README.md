Installation de Golang

Rendez-vous sur le site officiel de Go à l'adresse https://golang.org/dl/. Vous y
trouverez les différentes versions de Go disponibles. Choisissez la version qui
correspond à votre système d'exploitation (Windows, macOS, Linux, etc.) et à votre
architecture (32 bits ou 64 bits).

Windows :
Exécutez le programme d'installation que vous avez téléchargé.
L'installation va créer un dossier Go dans le répertoire spécifié (par défaut, C:\Go ).
Assurez-vous que le répertoire d'installation de Go est ajouté au chemin d'accès
système. Pour cela, accédez aux Propriétés du système > Paramètres système avancés >
Variables d'environnement > Path . Ajoutez le chemin vers le dossier Go (par exemple,
C:\Go\bin ).

MacOS :
Ouvrez le package d'installation (.pkg) que vous avez téléchargé.
Suivez les instructions de l'assistant d'installation pour installer Go dans le
répertoire par défaut ( /usr/local/go ).
Vous devez également ajouter le répertoire de Go à votre variable PATH. Ouvrez le
Terminal et éditez votre fichier de configuration de shell ( ~/.bash_profile , ~/.bashrc ,
ou ~/.zshrc ) en ajoutant la ligne suivante : export PATH=/usr/local/go/bin:$PATH

Linux :
Décompressez l'archive que vous avez téléchargée dans un répertoire de votre
choix (par exemple, /usr/local/go ).
Vous devez également ajouter le répertoire de Go à votre variable PATH. Pour cela,
éditez le fichier ~/.profile ou ~/.bash_profile (selon votre configuration) et ajoutez
la ligne suivante : export PATH=/usr/local/go/bin:$PATH

Vérification de l'installation :
Ouvrez un nouveau terminal (ou invite de commande) et tapez "go version". Si tout est
installé correctement, vous devriez voir la version de Go s'affiche

Credit
Guylan
Anderson
Maxime
Lucas 
