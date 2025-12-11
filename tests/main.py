import sys
import logging

class FrontendApp:
    def __init__(self):
        self.logger = logging.getLogger(__name__)
        self.logger.setLevel(logging.INFO)
        handler = logging.StreamHandler(sys.stdout)
        handler.setFormatter(logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s'))
        self.logger.addHandler(handler)

    def run(self):
        try:
            self.logger.info('Starting frontend app')
            # Initialize application components
            self.init_components()
            # Start application loop
            self.app_loop()
        except Exception as e:
            self.logger.error(f'Error: {str(e)}')
            sys.exit(1)

    def init_components(self):
        self.logger.info('Initializing application components')
        # Initialize component 1
        self.component1 = Component1()
        # Initialize component 2
        self.component2 = Component2()

    def app_loop(self):
        self.logger.info('Starting application loop')
        while True:
            # Handle user input
            user_input = input('Enter command: ')
            if user_input == 'exit':
                self.logger.info('Exiting application loop')
                break
            # Process user input
            self.process_input(user_input)

    def process_input(self, user_input):
        self.logger.info(f'Processing user input: {user_input}')
        # Process input using component 1
        result = self.component1.process(user_input)
        # Process result using component 2
        final_result = self.component2.process(result)
        self.logger.info(f'Final result: {final_result}')

class Component1:
    def process(self, input_data):
        # Simulate processing
        return input_data.upper()

class Component2:
    def process(self, input_data):
        # Simulate processing
        return input_data.lower()

if __name__ == '__main__':
    app = FrontendApp()
    app.run()