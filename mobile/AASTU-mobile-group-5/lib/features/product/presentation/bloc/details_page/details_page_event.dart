// details_page_event.dart
part of 'details_page_bloc.dart';

abstract class DetailsPageEvent extends Equatable {
  const DetailsPageEvent();

  @override
  List<Object> get props => [];
}

class FetchProductByIdEvent extends DetailsPageEvent {
  final GetProductParams params;

  const FetchProductByIdEvent(this.params);

  @override
  List<Object> get props => [params];
}

class DeleteDetailsEvent extends DetailsPageEvent {
  final DeleteProductParams params;

  const DeleteDetailsEvent(this.params);

  @override
  List<Object> get props => [params];
}
