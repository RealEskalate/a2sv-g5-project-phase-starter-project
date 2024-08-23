import 'package:flutter/material.dart';

class withTime extends StatelessWidget {
  const withTime({
    Key? key,
    required this.text,
    required this.image,
    required this.isCurrentUser,
    required this.type,
    required this.time,
  }): super(key: key);
  
  final String? text;
  final String? image;
  final String type;
  final bool isCurrentUser;
  final String time;

  @override
  Widget build(BuildContext context) {
    return ChatBubble(text: text, image: image, isCurrentUser: isCurrentUser, type: type, time: time);
  }
}

class ChatBubble extends StatelessWidget {
  const ChatBubble({
    Key? key,
    required this.text,
    required this.image,
    required this.isCurrentUser,
    required this.type,
    required this.time,
  }) : super(key: key);

  final String? text;
  final String? image;
  final String type;
  final bool isCurrentUser;
  final String time;

  @override
  Widget build(BuildContext context) {
    if(type == 'image') {
      if(image != null) {
        return Imagetype(isCurrentUser: isCurrentUser, image: image, time: time);
      } else {
        return const Placeholder();
      }
    } else if(type == 'text'){
      return TextType(isCurrentUser: isCurrentUser, text: text, time: time,);
    } else if (type == 'audio') {
      return VoiceMessageBubble(isCurrentUser: isCurrentUser, duration: '20');
    } else {
      return const Placeholder();
    }
  }
}

class Imagetype extends StatelessWidget {
  const Imagetype({
    super.key,
    required this.isCurrentUser,
    required this.image, 
    required this.time,
  });

  final bool isCurrentUser;
  final String? image;
  final String time;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.fromLTRB(
        isCurrentUser ? 64.0 : 16.0,
        4,
        isCurrentUser ? 16.0 : 64.0,
        4,
      ),
      child: Align(
        alignment: isCurrentUser ? Alignment.centerRight : Alignment.centerLeft,
        child: Column(
          crossAxisAlignment: !isCurrentUser? CrossAxisAlignment.end : CrossAxisAlignment.start,
          children: [
            ClipRRect(
              borderRadius:  const BorderRadius.all(Radius.circular(20)),
              child: Image.network(
                image!,
                fit: BoxFit.fill
              ),
            ),
            const SizedBox(height: 10,),
            Text(time)
          ],
        ),
      ),
    );
  }
}

class TextType extends StatelessWidget {
  const TextType({
    super.key,
    required this.isCurrentUser,
    required this.text,
    required this.time,
  });

  final bool isCurrentUser;
  final String? text;
  final String time;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.fromLTRB(
        isCurrentUser ? 64.0 : 16.0,
        4,
        isCurrentUser ? 16.0 : 64.0,
        4,
      ),
      child: Align(
        alignment: isCurrentUser ? Alignment.centerRight : Alignment.centerLeft,
        child: Column(
          crossAxisAlignment: !isCurrentUser? CrossAxisAlignment.end : CrossAxisAlignment.start,
          children: [
            ClipRect(
              child: DecoratedBox(
                decoration: BoxDecoration(
                  color: isCurrentUser ? Colors.blue : const Color(0XFFF2F7FB),
                  borderRadius: isCurrentUser? const BorderRadius.only(
                    topLeft: Radius.circular(20),
                    bottomLeft: Radius.circular(20),
                    bottomRight: Radius.circular(20),
                  ) : const BorderRadius.only(
                    topRight: Radius.circular(20),
                    bottomLeft: Radius.circular(20),
                    bottomRight: Radius.circular(20),
                  ) ,
                ),
                child: Padding(
                  padding: const EdgeInsets.all(20),
                  child: Text(
                    text!,
                    style: TextStyle(
                      fontSize: 17,
                      fontWeight: FontWeight.w500,
                      color: isCurrentUser ? Colors.white : Colors.black
                    ),
                  ),
                ),
              ),
            ),
            const SizedBox(height: 10,),
            Text(time),
          ],
        ),
      ),
    );
  }
}

class VoiceMessageBubble extends StatelessWidget {
  final bool isCurrentUser;
  final String duration;

  const VoiceMessageBubble({
    Key? key,
    required this.isCurrentUser,
    required this.duration,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.fromLTRB(
        isCurrentUser ? 64.0 : 16.0,
        4,
        isCurrentUser ? 16.0 : 64.0,
        4,
      ),
      child: Align(
        alignment: isCurrentUser ? Alignment.centerRight : Alignment.centerLeft,
        child: Container(
          padding: const EdgeInsets.all(10),
          decoration: BoxDecoration(
            color: isCurrentUser ? Colors.blue[50] : const Color(0xFFF2F7FB),
            borderRadius: BorderRadius.circular(20),
          ),
          child: Row(
            mainAxisSize: MainAxisSize.min,
            children: [
              // Play Button
              Container(
                width: 40,
                height: 40,
                decoration: const BoxDecoration(
                  shape: BoxShape.circle,
                  color: Colors.white,
                ),
                child: const Icon(
                  Icons.play_arrow,
                  color: Colors.purple,
                ),
              ),
              const SizedBox(width: 10),
              // Waveform
              Expanded(
                child: Container(
                  height: 40,
                  child: CustomPaint(
                    painter: WaveformPainter(),
                  ),
                ),
              ),
              const SizedBox(width: 10),
              // Duration
              Text(
                duration,
                style: const TextStyle(
                  color: Colors.black,
                  fontSize: 16,
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class WaveformPainter extends CustomPainter {
  @override
  void paint(Canvas canvas, Size size) {
    final paint = Paint()
      ..color = Colors.purple
      ..strokeCap = StrokeCap.round
      ..strokeWidth = 2;

    final path = Path();

    final waveform = [5, 10, 20, 15, 30, 40, 25, 10, 5]; // Sample waveform data

    final widthPerWave = size.width / (waveform.length * 2);

    for (int i = 0; i < waveform.length; i++) {
      final x = i * widthPerWave * 2;
      final y = size.height / 2;
      path.moveTo(x, y - waveform[i] / 2);
      path.lineTo(x, y + waveform[i] / 2);
    }

    canvas.drawPath(path, paint);
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return false;
  }
}
