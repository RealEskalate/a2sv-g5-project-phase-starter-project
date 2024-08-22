abstract class UpdateState {}



class UpdateLoading extends UpdateState {}

class UpdateSuccess extends UpdateState {

}

class UpdateFailiure extends UpdateState {
  final String error;

  UpdateFailiure({required this.error});
}
