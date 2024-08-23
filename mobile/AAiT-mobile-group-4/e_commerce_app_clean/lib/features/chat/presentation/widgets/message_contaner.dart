import 'dart:async';
import 'dart:io';

import 'package:audio_waveforms/audio_waveforms.dart';
import 'package:auto_size_text/auto_size_text.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class WithTime extends StatelessWidget {
  const WithTime({
    super.key,
    required this.text,
    required this.image,
    required this.isCurrentUser,
    required this.type,
    required this.time,
  });

  final String? text;
  final String? image;
  final String type;
  final bool isCurrentUser;
  final String time;

  @override
  Widget build(BuildContext context) {
    return ChatBubble(text: text, image: image, isCurrentUser: isCurrentUser, type: type, time: time,);
  }
}

class ChatBubble extends StatelessWidget {
  const ChatBubble({
    super.key,
    required this.text,
    required this.image,
    required this.isCurrentUser,
    required this.type,
    required this.time,
  });
  
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
        isCurrentUser ? 40.0 : 10.0,
        4,
        isCurrentUser ? 10.0 : 40.0,
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
                  padding: const EdgeInsets.all(10),
                  child: SizedBox(
                    width: MediaQuery.of(context).size.width - 150,
                    child: AutoSizeText(
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
            ),
            const SizedBox(height: 10,),
            Text(time),
          ],
        ),
      ),
    );
  }
}


class WaveBubble extends StatefulWidget {
  final bool isSender;
  final int? index;
  final String? path;
  final double? width;
  final Directory appDirectory;

  const WaveBubble({
    super.key,
    required this.appDirectory,
    this.width,
    this.index,
    this.isSender = false,
    this.path,
  });

  @override
  State<WaveBubble> createState() => _WaveBubbleState();
}

class _WaveBubbleState extends State<WaveBubble> {
  File? file;

  late PlayerController controller;
  late StreamSubscription<PlayerState> playerStateSubscription;

  final playerWaveStyle = const PlayerWaveStyle(
    fixedWaveColor: Colors.white54,
    liveWaveColor: Colors.white,
    spacing: 6,
  );

  @override
  void initState() {
    super.initState();
    controller = PlayerController();
    _preparePlayer();
    playerStateSubscription = controller.onPlayerStateChanged.listen((_) {
      setState(() {});
    });
  }

  void _preparePlayer() async {
    // Opening file from assets folder
    if (widget.index != null) {
      file = File('${widget.appDirectory.path}/audio${widget.index}.mp3');
      await file?.writeAsBytes(
          (await rootBundle.load('assets/audios/audio1.mp3')).buffer.asUint8List());
    }
    if (widget.index == null && widget.path == null && file?.path == null) {
      return;
    }
    // Prepare player with extracting waveform if index is even.
    controller.preparePlayer(
      path: widget.path ?? file!.path,
      shouldExtractWaveform: widget.index?.isEven ?? true,
    );
    // Extracting waveform separately if index is odd.
    if (widget.index?.isOdd ?? false) {
      controller
          .extractWaveformData(
            path: widget.path ?? file!.path,
            noOfSamples:
                playerWaveStyle.getSamplesForWidth(widget.width ?? 200),
          )
          .then((waveformData) => debugPrint(waveformData.toString()));
    }
  }

  @override
  void dispose() {
    playerStateSubscription.cancel();
    controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return widget.path != null || file?.path != null
        ? Align(
            alignment:
                widget.isSender ? Alignment.centerRight : Alignment.centerLeft,
            child: Container(
              padding: EdgeInsets.only(
                bottom: 6,
                right: widget.isSender ? 0 : 10,
                top: 6,
              ),
              margin: const EdgeInsets.symmetric(vertical: 8, horizontal: 12),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                color: widget.isSender
                    ? const Color(0xFF276bfd)
                    : const Color(0xFF343145),
              ),
              child: Row(
                mainAxisSize: MainAxisSize.min,
                children: [
                  if (!controller.playerState.isStopped)
                    IconButton(
                      onPressed: () async {
                        controller.playerState.isPlaying
                            ? await controller.pausePlayer()
                            : await controller.startPlayer(
                                finishMode: FinishMode.loop,
                              );
                      },
                      icon: Icon(
                        controller.playerState.isPlaying
                            ? Icons.stop
                            : Icons.play_arrow,
                      ),
                      color: Colors.white,
                      splashColor: Colors.transparent,
                      highlightColor: Colors.transparent,
                    ),
                  AudioFileWaveforms(
                    size: Size(MediaQuery.of(context).size.width / 2, 70),
                    playerController: controller,
                    waveformType: widget.index?.isOdd ?? false
                        ? WaveformType.fitWidth
                        : WaveformType.long,
                    playerWaveStyle: playerWaveStyle,
                  ),
                  if (widget.isSender) const SizedBox(width: 10),
                ],
              ),
            ),
          )
        : const SizedBox.shrink();
  }
}