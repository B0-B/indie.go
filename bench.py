import os 

# -- dependency check
# install prology for benchmark
#os.system('pip3 install git+https://github.com/B0-B/prology.git')

# load prology
from prology.log import logger
log = logger()

for size in [50, 250, 500, 750, 1052, 2000, 5000]:
    def test():
        os.system(f"go run indie.go -e -o ./originals/test_{size}.png -t ./targets/target_{size}.png")
    log.note(f"size {size} px", benchmark=test)