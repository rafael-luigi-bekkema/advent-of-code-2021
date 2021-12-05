;; -*- lexical-binding: t -*-

(defun my/day01atest ()
  (let* ((input "199\n200\n208\n210\n200\n207\n240\n269\n260\n263")
         (nums (mapcar 'string-to-number (split-string input "\n")))
         (expect 7)
         count)
    (setq count (my/day01acore nums))
    (cl-assert (eq expect count))
    count))

(defun my/day01acore (nums)
  (let* ((count 0)
         prev)
    (dolist (n nums count)
      (when (and prev (> n prev))
        (setq count (1+ count)))
      (setq prev n))
    count))

(defun my/day01a ()
  (let ((nums (mapcar 'string-to-number (process-lines "cat" "input/day01.txt"))))
    (my/day01acore nums)))
            
(progn
(my/day01test)

(my/day01a))
