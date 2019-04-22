class LocalStore:
    def __init__(self, base_dir):
        self.base_dir = base_dir

    def clean(self):
        import shutil
        try:
            shutil.rmtree(self.base_dir)
        except:
            pass
