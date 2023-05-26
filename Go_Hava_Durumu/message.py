import pywhatkit


def send_whatsapp_message(number, message, time_hour, time_min):
    pywhatkit.sendwhatmsg(number, message, time_hour, time_min)

def send_message_from_text_file(file_path, number, time_hour, time_min):
    with open(file_path, 'r') as file:
        message = file.read()
        send_whatsapp_message(number, message, time_hour, time_min)


file_path = 'weather_info.txt'
number = ''
time_hour = 14 
time_min = 24

send_message_from_text_file(file_path, number, time_hour, time_min)
