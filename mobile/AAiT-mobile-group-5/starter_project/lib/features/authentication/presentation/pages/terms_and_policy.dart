import 'package:flutter/material.dart';

class TermsAndPolicy extends StatelessWidget {
  const TermsAndPolicy({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Terms and Policy"),
        backgroundColor: Colors.white,
      ),
      body: Padding(
        padding: const EdgeInsets.symmetric(vertical: 16.0, horizontal: 25.0),
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text(
                "Terms and Conditions",
                style: TextStyle(
                  fontSize: 26,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF3F51F3),
                  letterSpacing: 1.2,
                ),
              ),
              const SizedBox(height: 20),
              const Text(
                "Welcome tro our application. By accessing or using our service, you agree to be bound by these terms and conditions. Please read them carefully.",
                style: TextStyle(fontSize: 16, height: 1.5),
              ),
              const SizedBox(height: 30),

              // Terms Sections
              _buildSectionTitle("1. Introduction"),
              _buildSectionContent(
                "These terms and conditions govern your use of this application. "
                "By using this application, you accept these terms and conditions in full. "
                "If you disagree with any part of these terms and conditions, do not use our application.",
              ),

              _buildSectionTitle("2. License to use the application"),
              _buildSectionContent(
                "Unless otherwise stated, we or our licensors own the intellectual property rights in the application and material on the application. "
                "Subject to the license below, all these intellectual property rights are reserved.",
              ),

              _buildSectionTitle("3. Acceptable use"),
              _buildSectionContent(
                "You must not use this application in any way that causes, or may cause, damage to the application or impairment of the availability or accessibility of the application.",
              ),

              _buildSectionTitle("4. User content"),
              _buildSectionContent(
                "In these terms and conditions, “your user content” means material (including without limitation text, images, audio material, video material, and audio-visual material) that you submit to this application.",
              ),

              _buildSectionTitle("5. No warranties"),
              _buildSectionContent(
                "This application is provided “as is” without any representations or warranties, express or implied. "
                "We make no representations or warranties in relation to this application or the information and materials provided on this application.",
              ),

              _buildSectionTitle("6. Limitation of liability"),
              _buildSectionContent(
                "We will not be liable to you in relation to the contents of, or use of, or otherwise in connection with, this application for any direct, indirect, special, or consequential loss.",
              ),

              _buildSectionTitle("7. Indemnity"),
              _buildSectionContent(
                "You hereby indemnify us and undertake to keep us indemnified against any losses, damages, costs, liabilities, and expenses (including without limitation legal expenses and any amounts paid by us to a third party in settlement of a claim or dispute) incurred or suffered by us arising out of any breach by you of any provision of these terms and conditions.",
              ),

              _buildSectionTitle("8. Breaches of these terms and conditions"),
              _buildSectionContent(
                "Without prejudice to our other rights under these terms and conditions, if you breach these terms and conditions in any way, we may take such action as we deem appropriate to deal with the breach, including suspending your access to the application, prohibiting you from accessing the application, blocking computers using your IP address from accessing the application, contacting your internet service provider to request that they block your access to the application, and/or bringing court proceedings against you.",
              ),

              _buildSectionTitle("9. Variation"),
              _buildSectionContent(
                "We may revise these terms and conditions from time to time. Revised terms and conditions will apply to the use of this application from the date of the publication of the revised terms and conditions on this application.",
              ),

              _buildSectionTitle("10. Assignment"),
              _buildSectionContent(
                "We may transfer, sub-contract, or otherwise deal with our rights and/or obligations under these terms and conditions without notifying you or obtaining your consent.",
              ),

              _buildSectionTitle("11. Severability"),
              _buildSectionContent(
                "If a provision of these terms and conditions is determined by any court or other competent authority to be unlawful and/or unenforceable, the other provisions will continue in effect.",
              ),

              _buildSectionTitle("12. Entire agreement"),
              _buildSectionContent(
                "These terms and conditions, together with our privacy policy, constitute the entire agreement between you and us in relation to your use of this application, and supersede all previous agreements in respect of your use of this application.",
              ),

              _buildSectionTitle("13. Law and jurisdiction"),
              _buildSectionContent(
                "These terms and conditions will be governed by and construed in accordance with the laws of [Your Country], and any disputes relating to these terms and conditions will be subject to the exclusive jurisdiction of the courts of [Your Country].",
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildSectionTitle(String title) {
    return Padding(
      padding: const EdgeInsets.only(top: 20.0, bottom: 10.0),
      child: Text(
        title,
        style: const TextStyle(
          fontSize: 20,
          fontWeight: FontWeight.bold,
          color: Color(0xFF3F51F3),
          letterSpacing: 1.2,
        ),
      ),
    );
  }

  Widget _buildSectionContent(String content) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 20.0),
      child: Text(
        content,
        style: const TextStyle(fontSize: 16, height: 1.5),
      ),
    );
  }
}
