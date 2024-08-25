part of 'input_validation_cubit.dart';

// ignore: must_be_immutable
sealed class InputValidationState extends Equatable {
  final bool name;
  final bool catagory;
  final bool price;
  final File? imageUrl;
  String? nameMessage;
  String? catagoryMessage;
  String? priceMessage;
  static final File newFile = File('');
  InputValidationState(
      {required this.name,
      required this.catagory,
      required this.price,
      this.imageUrl,
      this.nameMessage,
      this.catagoryMessage,
      this.priceMessage});

  @override
  List<Object> get props => [
        name,
        catagory,
        price,
        imageUrl ?? newFile,
      ];
}

// ignore: must_be_immutable
final class InputValidationInitial extends InputValidationState {
  InputValidationInitial()
      : super(
          name: true,
          catagory: true,
          price: true,
        );
}

// ignore: must_be_immutable
final class InputValidatedState extends InputValidationState {
  @override
  // ignore: overridden_fields
  final bool name;
  @override
  // ignore: overridden_fields
  final bool catagory;
  @override
  // ignore: overridden_fields
  final bool price;
  @override
  // ignore: overridden_fields
  String? nameMessage;
  @override
  // ignore: overridden_fields
  String? catagoryMessage;
  @override
  // ignore: overridden_fields
  String? priceMessage;
  @override
  // ignore: overridden_fields
  File? imageUrl;
  InputValidatedState(
      {required this.name,
      required this.catagory,
      required this.price,
      this.imageUrl,
      this.nameMessage,
      this.catagoryMessage,
      this.priceMessage})
      : super(name: name, catagory: catagory, price: price, imageUrl: imageUrl);
  List<dynamic> getState(String type) {
    Map<String, List<dynamic>> correspond = {
      'Name': [name, nameMessage],
      'Price': [price, priceMessage],
      'Catagory': [catagory, catagoryMessage]
    };
    if (!correspond.containsKey(type)) {
      return [true, ''];
    } else {
      return correspond[type]!;
    }
  }

  bool isValidated() {
    bool value = name & price & catagory & (imageUrl != null);
    return value;
  }

  bool isValidForUpdate() {
    bool value = name && price && catagory;
    return value;
  }
}
